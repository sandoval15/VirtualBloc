import subprocess

def main():
    subprocess.Popen("uv run watcher.py", text=True)
    runServer("go run", "backend/server/local-server.go")
    runServer("node", "frontend/server/local-server.js")

def runServer(cmd: str, dir:str):
    cwd = dir.split("/", 1)
    subprocess.Popen(f"{cmd} {cwd[1]}", cwd=cwd[0], text=True)
    print(f"\nIniciando SERVIDOR de {cwd[0]}: \n")

if __name__ == "__main__":
    main()
