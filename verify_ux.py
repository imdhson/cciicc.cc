from playwright.sync_api import sync_playwright

def verify():
    with sync_playwright() as p:
        browser = p.chromium.launch()
        page = browser.new_page()
        page.goto("http://localhost:8080/host")
        page.screenshot(path="before.png")
        browser.close()

if __name__ == "__main__":
    verify()
