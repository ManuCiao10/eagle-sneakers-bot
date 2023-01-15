from dotenv import load_dotenv
import threading
from dsm.start import Start
import sys
from db import db

load_dotenv()


def main():
    threading.Thread(target=Start).start()

    while True:
        q = input()
        if q.lower() in ["q", "quit"]:
            db.close()
            sys.exit(0)


if __name__ == "__main__":
    main()
