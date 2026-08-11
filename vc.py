import sys
import subprocess

def main():
    if len(sys.argv) > 1:
        if sys.argv[1] == "run":
            sv = dict(
                back=runServer("go run", "backend/server/local-server.go"),
                front=runServer("node", "frontend/server/local-server.js")
            )
            while True:
                action = ""
                action = input("action: ").lower()

                match action:   
                    case "exit":
                        return
                    case _:
                        continue

def runServer(cmd: str, dir:str):
    cwd = dir.split("/", 1)
    sv = subprocess.Popen([{cmd}, {cwd[1]}], cwd=cwd[0], stdout=subprocess.PIPE, text=True)
    print(f"cmd: {cmd} {cwd[1]}")
    print(f"\nIniciando SERVIDOR de {cwd[0]}: \n")
    for line in sv.stdout:
        l = line.strip()
        print(l)
        if "Servidor corriendo" in l:
            break
    return sv

if __name__ == "__main__":
    main()
