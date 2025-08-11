import requests
import re

def get_latest_rufus():
    url = "https://api.github.com/repos/aristocratos/btop4win/releases/latest"
    response = requests.get(url)
    
    if response.status_code == 200:
        release_data = response.json()
        
        # Find the .exe asset
        for asset in release_data['assets']:
            if asset['name'].endswith('.zip'):
                return {
                    'version': release_data['tag_name'],
                    'download_url': asset['browser_download_url'],
                    'name': asset['name']
                }
    
    return None

latest = get_latest_rufus()
if latest:
    print(f"Latest version: {latest['version']}")
    print(f"Download URL: {latest['download_url']}")