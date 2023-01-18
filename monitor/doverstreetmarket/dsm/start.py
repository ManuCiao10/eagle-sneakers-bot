from dsm.types import Thread
from dsm.backendlink import BackendLinkFlow
from db import db
from dsm.keyword import KeywordFlow


def Start():
    t = Thread(BackendLinkFlow, None)
    t.start()

    # cur = db.cursor()

    # keywordRows = cur.execute("SELECT keyword FROM keywords").fetchall()
    # keywords = [keyword[0] for keyword in keywordRows]
    keywordThread = Thread(KeywordFlow, t)
    keywordThread.start()
