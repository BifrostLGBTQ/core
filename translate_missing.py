#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import json
import time
import sys

try:
    from deep_translator import GoogleTranslator
except ImportError:
    print("deep-translator yükleniyor...")
    import subprocess
    subprocess.check_call([sys.executable, "-m", "pip", "install", "deep-translator", "--quiet", "--user"])
    from deep_translator import GoogleTranslator

# JSON dosyasını oku
with open('static/data/sexual_preferences.json', 'r', encoding='utf-8') as f:
    data = json.load(f)

language_map = {
    'tr': 'tr', 'es': 'es', 'he': 'iw', 'ar': 'ar', 'zh': 'zh-CN', 
    'ja': 'ja', 'hi': 'hi', 'de': 'de', 'th': 'th', 'ru': 'ru', 
    'pl': 'pl', 'fr': 'fr', 'pt': 'pt', 'id': 'id', 'bn': 'bn'
}

languages = list(language_map.keys())

# Çevirisi eksik olan item'ları bul
items_to_translate = []
for idx, item in enumerate(data):
    needs_translation = False
    
    # Label'larda boş olanları kontrol et
    for lang in languages:
        if not item['label'].get(lang) or item['label'][lang] == "":
            needs_translation = True
            break
    
    # Description'larda boş olanları kontrol et
    if not needs_translation:
        for lang in languages:
            if not item['description'].get(lang) or item['description'][lang] == "":
                needs_translation = True
                break
    
    if needs_translation:
        items_to_translate.append((idx, item))

total = len(items_to_translate)
print(f"Toplam {total} item çevirisi eksik.\n")

# Tüm item'ları çevir
for batch_idx, (idx, item) in enumerate(items_to_translate):
    en_label = item['label']['en']
    en_desc = item['description']['en']
    
    print(f"[{batch_idx+1}/{total}] {en_label}...")
    
    # Label çevirileri
    for lang_code, google_code in language_map.items():
        if not item['label'].get(lang_code) or item['label'][lang_code] == "":
            try:
                translator = GoogleTranslator(source='en', target=google_code)
                translated = translator.translate(en_label)
                item['label'][lang_code] = translated
                time.sleep(0.2)
            except Exception as e:
                print(f"  ⚠ {lang_code} label: {str(e)[:50]}")
                # Retry once
                try:
                    time.sleep(1)
                    translator = GoogleTranslator(source='en', target=google_code)
                    translated = translator.translate(en_label)
                    item['label'][lang_code] = translated
                except:
                    pass
    
    # Description çevirileri
    for lang_code, google_code in language_map.items():
        if not item['description'].get(lang_code) or item['description'][lang_code] == "":
            try:
                translator = GoogleTranslator(source='en', target=google_code)
                translated = translator.translate(en_desc)
                item['description'][lang_code] = translated
                time.sleep(0.2)
            except Exception as e:
                print(f"  ⚠ {lang_code} desc: {str(e)[:50]}")
                # Retry once
                try:
                    time.sleep(1)
                    translator = GoogleTranslator(source='en', target=google_code)
                    translated = translator.translate(en_desc)
                    item['description'][lang_code] = translated
                except:
                    pass
    
    # Her 5 item'da bir kaydet
    if (batch_idx + 1) % 5 == 0:
        with open('static/data/sexual_preferences.json', 'w', encoding='utf-8') as f:
            json.dump(data, f, ensure_ascii=False, indent=4)
        print(f"  ✓ Kaydedildi ({batch_idx+1}/{total})\n")
        time.sleep(2)  # Rate limiting

# Final kayıt
with open('static/data/sexual_preferences.json', 'w', encoding='utf-8') as f:
    json.dump(data, f, ensure_ascii=False, indent=4)

print(f"\n✓ Tüm çeviriler tamamlandı! {total} item çevrildi.")

