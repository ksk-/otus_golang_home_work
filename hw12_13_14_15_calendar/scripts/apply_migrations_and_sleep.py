import asyncio
import subprocess
import sys

SLEEP_TIMEOUT_SEC = 5


async def main():
    args = ['/usr/local/bin/goose'] + sys.argv[1:]
    process = subprocess.Popen(args, stdout=subprocess.PIPE)
    for line in process.communicate():
        print(line)

    with open('ready', 'w') as f:
        pass

    while True:
        await asyncio.sleep(SLEEP_TIMEOUT_SEC)

if __name__ == '__main__':
    try:
        asyncio.run(main())
    except KeyboardInterrupt:
        pass
