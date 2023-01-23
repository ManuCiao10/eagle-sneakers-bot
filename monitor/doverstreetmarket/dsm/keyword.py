from dotenv import load_dotenv
from dsm.types import Thread, Size
import requests

from dsm.webhook import PingFrontend
import time

load_dotenv()


def KeywordFlow(keywords: list, parentThread: Thread):
    print("[KEYWORDS] Starting thread.")
    site = "doverstreetmarket"
    keywords = ["dunk", "jordan", "retro", "yeezy", "off-white", "nike", "balance"]

    headers = {
        "authority": "shop.doverstreetmarket.com",
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
        "accept-language": "en-GB,en;q=0.7",
        "cache-control": "max-age=0",
        "sec-ch-ua": '"Not?A_Brand";v="8", "Chromium";v="108", "Brave";v="108"',
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": '"macOS"',
        "sec-fetch-dest": "document",
        "sec-fetch-mode": "navigate",
        "sec-fetch-site": "none",
        "sec-fetch-user": "?1",
        "sec-gpc": "1",
        "upgrade-insecure-requests": "1",
        "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
    }
    previousItems = []

    while not parentThread.stop:
        time.sleep(1)
        try:
            print("[BACKENDLINK] Checking for new items...")
            try:
                response = requests.get(
                    "https://shop.doverstreetmarket.com/products.json", headers=headers
                )
            except requests.exceptions.ConnectionError or requests.exceptions.ConnectTimeout:
                print("[KEYWORD] Connection error.")
                continue
            if response.status_code == 200:
                print("[BACKENDLINK] Got response...")
                data = response.json()
                foundProduct = False

                for product in data["products"]:
                    item_name = product["title"]
                    pid = product["handle"].upper()
                    price = product["variants"][0]["price"]

                    if parentThread.firstRun:
                        print(f"[KEYWORDS] Adding firstrun product: {pid}")
                        previousItems.append(pid)
                        continue
                    else:
                        if pid not in previousItems and any(
                            keyword in item_name.lower().split(" ")
                            for keyword in keywords
                        ):
                            previousItems.append(pid)
                            foundProduct = True
                            print(f"[KEYWORDS] Found new product: {pid}")

                            for image in product["images"]:
                                image = image["src"]
                                break

                            variants = []
                            url = f"https://shop.doverstreetmarket.com/products/{pid}.json"

                            for variant in product["variants"]:
                                if variant["available"] == True:
                                    var = variant["title"]

                                    variants.append(var)

                            PingFrontend(
                                webhookImage=image,
                                productSizes=variants,
                                title=item_name,
                                price=price,
                                url=url,
                                pid=pid,
                                site=site.upper(),
                            )

                if parentThread.firstRun:
                    parentThread.firstRun = False

                if not foundProduct:
                    print("[KEYWORDS] No new products found.")

            else:
                print("[BACKENDLINK] Got error response...", response.status_code)
                time.sleep(1)
                continue

        except Exception as e:
            print("[KEYWORD] Error: {}".format(e))
            continue
