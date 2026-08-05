import http from 'node:http';
import fs from 'node:fs';
import { transpileToRam } from './transpiler.js';
import path from 'path';
import { fileURLToPath } from 'url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));

const PORT = 3000
const HOST = 'localhost'

let js = await transpileToRam()

const server = http.createServer((req, res) => {
    const { url, method } = req
    if (url === '/' && method === 'GET') {
        res.writeHead(200, { 'Content-Type': 'text/html; charset=utf-8' });
        let html = fs.readFileSync(path.join(__dirname, '../main.html'), 'utf-8')
        res.end(html);
        return;
    }

    if (url.endsWith('.css') && method === 'GET') {
        res.writeHead(200, { 'Content-Type': 'text/css; charset=utf-8' });
        let dir = path.join(__dirname, `..${url}`)
        let css = fs.readFileSync(dir, 'utf-8')
        res.end(css);
        return;
    }

    if (url.endsWith('.js') && method === 'GET') {
        res.writeHead(200, { 'Content-Type': 'text/javascript; charset=utf-8' });
        let dir = path.join(__dirname, `..${url}`)
        console.log('Requesting JS file: ', dir)
        if (fs.existsSync(dir)) {
            let code = fs.readFileSync(dir, 'utf-8')
            res.end(code);
            return;
        }
        let code = js.find(f => dir === f.path).code
        res.end(Buffer.from(code));
        return;
    }
})

server.listen(PORT, HOST, () => {
    console.log(`Servidor corriendo en http://${HOST}:${PORT}`)
})