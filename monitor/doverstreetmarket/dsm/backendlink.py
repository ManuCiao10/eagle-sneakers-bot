from dsm.types import Thread, Size
import requests
import time
import xmltodict
from dsm.webhook import PingServer


def GetPIDFromLink(link: str) -> str:
    return link.split("/")[-1].upper()


def BackendLinkFlow(_, parentThread: Thread):
    print("[BACKENDLINK] Starting thread.")

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

    params = {
        "from": "6216894087324",
        "to": "7152017375494",
    }

    links = []
    site = "doverstreetmarket"

    while not parentThread.stop:
        try:
            time.sleep(2)
            try:
                response = requests.get(
                    "https://shop.doverstreetmarket.com/sitemap_products_1.xml",
                    params=params,
                    headers=headers,
                )
            except requests.exceptions.ConnectionError or requests.exceptions.ConnectTimeout:
                print("[BACKENDLINK] Connection error.")
                continue

            if response.status_code == 200:
                try:
                    data = xmltodict.parse(response.text)
                except xmltodict.expat.ExpatError:
                    print("[BACKENDLINK] XML Error.")
                    continue

                print("[BACKENDLINK] Successfully fetched data.")
                if not data["urlset"].get("url"):
                    print("[BACKENDLINK] No links found.")
                    continue

                if parentThread.firstRun:
                    parentThread.firstRun = False

                    for url in data["urlset"]["url"]:
                        links.append(url["loc"])
                        print("[BACKENDLINK] Added link: {}".format(url["loc"]))

                    continue

                if [url["loc"] for url in data["urlset"]["url"]] == links:
                    print("[BACKENDLINK] No new links found.")
                    continue

                for url in data["urlset"]["url"]:
                    if url["loc"] not in links:
                        links.append(url["loc"])
                        print("[BACKENDLINK] Added link: {}".format(url["loc"]))

                        while True:
                            sizes = []
                            try:
                                response = requests.get(
                                    url["loc"] + ".json", headers=headers
                                )

                            except requests.exceptions.ConnectionError or requests.exceptions.ConnectTimeout:
                                print("[BACKENDLINK] Connection error.")

                            print("[RESPONSE] {}".format(response.status_code))
                            if response.status_code != 200:
                                time.sleep(1)
                                continue

                            pid = GetPIDFromLink(url["loc"])

                            json = response.json()
                            title = json["product"]["title"]
                            image = json["product"]["image"]["src"]
                            price = json["product"]["variants"][0]["price"]
                            for size in json["product"]["options"]:
                                for value in size["values"]:
                                    sizes.append(value)
                            # sizeMqt = json

                            PingServer(
                                webhookImage=image,
                                productSizes=sizes,
                                title=title,
                                price=price,
                                url=url["loc"],
                                pid=pid,
                                site=site.upper(),
                            )

                            break
                links = [url["loc"] for url in data["urlset"]["url"]]
            else:
                print("[BACKENDLINK] Error fetching data.")
                time.sleep(3)

        except Exception as e:
            print("[BACKENDLINK] Error: {}".format(e))
            continue
