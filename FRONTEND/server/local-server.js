import http from 'node:http';
import fs from 'node:fs';
import { transpileToRam, reTranspile } from './transpiler.js';
import path from 'path';
import { fileURLToPath } from 'url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));

const PORT = 3000
const HOST = 'localhost'

const notFound = (res) => {
    res.writeHead(404, { 'Content-Type': 'text/plain; charset=utf-8' });
    res.end('404 - No encontrado')
}

let js = await transpileToRam()

const server = http.createServer((req, res) => {
    const { url, method } = req
    if (method === 'GET') {
        if (url.split('?')[0] === '/retranspile') {
            res.writeHead(200, { 'Content-Type': 'charset=utf-8' });
            let file = new URL(url, "http://localhost:3000").searchParams.get("file").toString()
            let dir = path.join(__dirname, '../js')
            js.forEach(f => {
                if (f.path === path.join(dir, file.replace('.ts', '.js'))) {
                    f.code = Buffer.from(reTranspile(file))
                }
            })
            res.end("ok")
            return
        }
        if (url === '/') {
            res.writeHead(200, { 'Content-Type': 'text/html; charset=utf-8' });
            let html = fs.readFileSync(path.join(__dirname, '../main.html'), 'utf-8')
            res.end(html)
            return
        }

        if (url.endsWith('.css')) {
            let dir = path.join(__dirname, `..${url}`)
            if (!fs.existsSync(dir)) {
                console.log('CSS file not found: ', dir)
                notFound(res)
                return
            }
            res.writeHead(200, { 'Content-Type': 'text/css; charset=utf-8' });
            let css = fs.readFileSync(dir, 'utf-8')
            res.end(css)
            return
        }

        if (url.endsWith('.js')) {
            let dir = path.join(__dirname, `..${url}`)
            console.log('Requesting JS file: ', dir)
            if (fs.existsSync(dir)) {
                res.writeHead(200, { 'Content-Type': 'text/javascript; charset=utf-8' });
                let code = fs.readFileSync(dir, 'utf-8')
                res.end(code)
                return
            }
            let file = js.find(f => dir === f.path)
            if (!file) {
                console.log('JS file not found on disk or in RAM: ', dir)
                notFound(res)
                return
            }
            res.writeHead(200, { 'Content-Type': 'text/javascript; charset=utf-8' });
            res.end(Buffer.from(file.code))
            return
        }
    }

    notFound(res)
})

server.listen(PORT, HOST, () => {
    console.log(`Servidor corriendo en http://${HOST}:${PORT}`)
})
