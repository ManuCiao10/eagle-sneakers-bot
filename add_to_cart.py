import requests

cookies = {
    "mage-translation-storage": "%7B%7D",
    "mage-translation-file-version": "%7B%7D",
    "_gcl_au": "1.1.434003817.1665959414",
    "rmStore": "ald:20220924_1801|atrv:nmrHekKy67Q-4dg2BmmR8wQ5hCXCLzqi6Q",
    "_ga": "GA1.1.133822203.1665959415",
    "CookieConsent": "{stamp:%27-1%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1665959415837%2Ciab2:%27%27%2Cregion:%27CA%27}",
    "sugar_newsletter": "1",
    "_hjSessionUser_2226440": "eyJpZCI6ImU5MTk5OGIzLTE2M2YtNTdmNC05MzA2LTA2NTE3N2ZmMzQzZiIsImNyZWF0ZWQiOjE2NjU5NTk0MTUxNTYsImV4aXN0aW5nIjp0cnVlfQ==",
    "mage-cache-storage": "%7B%7D",
    "mage-cache-storage-section-invalidation": "%7B%7D",
    "mage-messages": "",
    "recently_viewed_product": "%7B%7D",
    "recently_viewed_product_previous": "%7B%7D",
    "recently_compared_product": "%7B%7D",
    "recently_compared_product_previous": "%7B%7D",
    "product_data_storage": "%7B%7D",
    "private_content_version": "e6593975982f0e43dcfeb8494bbd97aa",
    "__stripe_mid": "4ed876e4-fdab-40de-b131-1f30ed73cadcc31d27",
    "_clck": "1vh2vhs|1|f5s|0",
    "PHPSESSID": "f5ed036ef7ab8a0a368f17465dc1a613",
    "X-Magento-Vary": "c58cc7336841735bf5ef13185766282824a9d073",
    "_hjIncludedInSessionSample": "0",
    "_hjSession_2226440": "eyJpZCI6IjM3ZGZlNTJmLWYyYzMtNDExYi05YTM4LTcxYzEzNGFjNmY0NCIsImNyZWF0ZWQiOjE2NjYwMDY3MTM0NDMsImluU2FtcGxlIjpmYWxzZX0=",
    "_hjIncludedInPageviewSample": "1",
    "_hjAbsoluteSessionInProgress": "0",
    "form_key": "dXGNPdDTbKeBPyNH",
    "mage-cache-sessid": "true",
    "section_data_ids": "%7B%22customer%22%3A1666006714%2C%22compare-products%22%3A1666006714%2C%22last-ordered-items%22%3A1666006714%2C%22cart%22%3A1666006715%2C%22directory-data%22%3A1666006715%2C%22review%22%3A1666006714%2C%22instant-purchase%22%3A1666006714%2C%22persistent%22%3A1666006714%2C%22captcha%22%3A1666006714%2C%22wishlist%22%3A1666006809%2C%22recently_viewed_product%22%3A1666006714%2C%22recently_compared_product%22%3A1666006714%2C%22product_data_storage%22%3A1666006714%2C%22paypal-billing-agreement%22%3A1666006714%2C%22checkout-fields%22%3A1666006714%2C%22collection-point-result%22%3A1666006714%2C%22pickup-location-result%22%3A1666006714%7D",
    "_clsk": "24z7l0|1666006823031|7|1|h.clarity.ms/collect",
    "_ga_1TT1ERKS8Z": "GS1.1.1666006712.2.1.1666006845.60.0.0",
}

headers = {
    "authority": "www.sugar.it",
    "accept": "application/json, text/javascript, */*; q=0.01",
    "accept-language": "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5",
    "cache-control": "no-cache",
    "content-type": "multipart/form-data; boundary=----WebKitFormBoundarymNRAWRTRzC0JSa6A",
    # Requests sorts cookies= alphabetically
    # 'cookie': 'mage-translation-storage=%7B%7D; mage-translation-file-version=%7B%7D; _gcl_au=1.1.434003817.1665959414; rmStore=ald:20220924_1801|atrv:nmrHekKy67Q-4dg2BmmR8wQ5hCXCLzqi6Q; _ga=GA1.1.133822203.1665959415; CookieConsent={stamp:%27-1%27%2Cnecessary:true%2Cpreferences:true%2Cstatistics:true%2Cmarketing:true%2Cver:1%2Cutc:1665959415837%2Ciab2:%27%27%2Cregion:%27CA%27}; sugar_newsletter=1; _hjSessionUser_2226440=eyJpZCI6ImU5MTk5OGIzLTE2M2YtNTdmNC05MzA2LTA2NTE3N2ZmMzQzZiIsImNyZWF0ZWQiOjE2NjU5NTk0MTUxNTYsImV4aXN0aW5nIjp0cnVlfQ==; mage-cache-storage=%7B%7D; mage-cache-storage-section-invalidation=%7B%7D; mage-messages=; recently_viewed_product=%7B%7D; recently_viewed_product_previous=%7B%7D; recently_compared_product=%7B%7D; recently_compared_product_previous=%7B%7D; product_data_storage=%7B%7D; private_content_version=e6593975982f0e43dcfeb8494bbd97aa; __stripe_mid=4ed876e4-fdab-40de-b131-1f30ed73cadcc31d27; _clck=1vh2vhs|1|f5s|0; PHPSESSID=f5ed036ef7ab8a0a368f17465dc1a613; X-Magento-Vary=c58cc7336841735bf5ef13185766282824a9d073; _hjIncludedInSessionSample=0; _hjSession_2226440=eyJpZCI6IjM3ZGZlNTJmLWYyYzMtNDExYi05YTM4LTcxYzEzNGFjNmY0NCIsImNyZWF0ZWQiOjE2NjYwMDY3MTM0NDMsImluU2FtcGxlIjpmYWxzZX0=; _hjIncludedInPageviewSample=1; _hjAbsoluteSessionInProgress=0; form_key=dXGNPdDTbKeBPyNH; mage-cache-sessid=true; section_data_ids=%7B%22customer%22%3A1666006714%2C%22compare-products%22%3A1666006714%2C%22last-ordered-items%22%3A1666006714%2C%22cart%22%3A1666006715%2C%22directory-data%22%3A1666006715%2C%22review%22%3A1666006714%2C%22instant-purchase%22%3A1666006714%2C%22persistent%22%3A1666006714%2C%22captcha%22%3A1666006714%2C%22wishlist%22%3A1666006809%2C%22recently_viewed_product%22%3A1666006714%2C%22recently_compared_product%22%3A1666006714%2C%22product_data_storage%22%3A1666006714%2C%22paypal-billing-agreement%22%3A1666006714%2C%22checkout-fields%22%3A1666006714%2C%22collection-point-result%22%3A1666006714%2C%22pickup-location-result%22%3A1666006714%7D; _clsk=24z7l0|1666006823031|7|1|h.clarity.ms/collect; _ga_1TT1ERKS8Z=GS1.1.1666006712.2.1.1666006845.60.0.0',
    "origin": "https://www.sugar.it",
    "pragma": "no-cache",
    "referer": "https://www.sugar.it/catalog/product/view/id/250253/s/gx1656-jade-green-forest-green/category/48/",
    "sec-ch-ua": '"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"',
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": '"macOS"',
    "sec-fetch-dest": "empty",
    "sec-fetch-mode": "cors",
    "sec-fetch-site": "same-origin",
    "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
    "x-requested-with": "XMLHttpRequest",
}

data = '------WebKitFormBoundarymNRAWRTRzC0JSa6A\r\nContent-Disposition: form-data; name="product"\r\n\r\n250253\r\n------WebKitFormBoundarymNRAWRTRzC0JSa6A\r\nContent-Disposition: form-data; name="selected_configurable_option"\r\n\r\n250244\r\n------WebKitFormBoundarymNRAWRTRzC0JSa6A\r\nContent-Disposition: form-data; name="related_product"\r\n\r\n\r\n------WebKitFormBoundarymNRAWRTRzC0JSa6A\r\nContent-Disposition: form-data; name="item"\r\n\r\n250253\r\n------WebKitFormBoundarymNRAWRTRzC0JSa6A\r\nContent-Disposition: form-data; name="form_key"\r\n\r\ndXGNPdDTbKeBPyNH\r\n------WebKitFormBoundarymNRAWRTRzC0JSa6A\r\nContent-Disposition: form-data; name="super_attribute[150]"\r\n\r\n40\r\n------WebKitFormBoundarymNRAWRTRzC0JSa6A--\r\n'

response = requests.post(
    "https://www.sugar.it/checkout/cart/add/uenc/aHR0cHM6Ly93d3cuc3VnYXIuaXQvY2F0YWxvZy9wcm9kdWN0L3ZpZXcvaWQvMjUwMjUzL3MvZ3gxNjU2LWphZGUtZ3JlZW4tZm9yZXN0LWdyZWVuL2NhdGVnb3J5LzQ4Lw%2C%2C/product/250253/",
    cookies=cookies,
    headers=headers,
    data=data,
)
