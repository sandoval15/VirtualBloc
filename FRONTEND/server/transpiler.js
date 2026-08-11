import esbuild from 'esbuild';
import fs from 'node:fs/promises';
import path from 'path';
import { fileURLToPath } from 'url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const dir = path.join(__dirname, '../ts')

export async function transpileToRam() {
    try {
        let files = await fs.readdir(dir, { withFileTypes: true, recursive: true })
        let ts = files.filter(f => f.isFile()).map(f => path.join(f.parentPath, f.name))
        console.log('Transpiling files: ', ts, '\n')

        let js = await esbuild.build({
            entryPoints: ts,
            bundle: true,
            write: false,
            outdir: path.join(__dirname, '../js'),
            format: 'esm',
            splitting: true,
            loader: {'.ts': 'ts'},
            minify: true,
            sourcemap: 'inline'
        })

        for (let f of js.outputFiles) {
            console.log(f.path +': '+ f.text.split("//#")[0] +'\n')
        }

        return js.outputFiles.map(file => ({ path: file.path, code: file.contents }))
    } catch (error) {
        console.log(error)
    }
}