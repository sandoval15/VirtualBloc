import ctypes
from ctypes import wintypes
import os
import urllib.request
import urllib.error
import urllib.parse

FILE_LIST_DIRECTORY = 0x0001
FILE_SHARE_READ = 0x00000001
FILE_SHARE_WRITE = 0x00000002
FILE_SHARE_DELETE = 0x00000004
OPEN_EXISTING = 3
FILE_FLAG_BACKUP_SEMANTICS = 0x02000000
FILE_NOTIFY_CHANGE_LAST_WRITE = 0x00000010

FILE_ACTION_MODIFIED = 3

kernel32 = ctypes.WinDLL("kernel32", use_last_error=True)
dir = os.path.abspath("frontend/ts")

handle = kernel32.CreateFileW(
    dir,
    FILE_LIST_DIRECTORY,
    FILE_SHARE_READ | FILE_SHARE_WRITE | FILE_SHARE_DELETE,
    None,
    OPEN_EXISTING,
    FILE_FLAG_BACKUP_SEMANTICS,
    None
)

if handle == -1:
    raise ctypes.WinError(ctypes.get_last_error())

size = 1024
buffer = ctypes.create_string_buffer(size)

while True:
    ok = kernel32.ReadDirectoryChangesW(
        handle,
        buffer,
        size,
        False,
        FILE_NOTIFY_CHANGE_LAST_WRITE,
        ctypes.byref(wintypes.DWORD()),
        None,
        None
    )

    if not ok:
        raise ctypes.WinError(ctypes.get_last_error())

    offset = 0
    while True:
        next_offset, action, name_len = struct_unpack = ctypes.cast(
            ctypes.byref(buffer, offset), ctypes.POINTER(wintypes.DWORD * 3)
        ).contents

        raw = buffer.raw[offset + 12: offset + 12 + name_len]
        file = raw.decode("utf-16-le")

        if file.split(".")[1] == "ts" and action == FILE_ACTION_MODIFIED:
            url = f"http://localhost:3000/retranspile?file={urllib.parse.quote(file)}"
            try:
                with urllib.request.urlopen(url, timeout=5) as res:
                    pass
            except urllib.error.URLError as e:
                print(f"No se pudo contactar al servidor: {e.reason}")

        if next_offset == 0:
            break
        offset += next_offset