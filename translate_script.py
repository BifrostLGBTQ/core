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
    subprocess.check_call([sys.executable, "-m", "pip", "install", "deep-translator", "--quiet"])
    from deep_translator import GoogleTranslator

# JSON dosyasını oku
with open('static/data/sexual_preferences.json', 'r', encoding='utf-8') as f:
    data = json.load(f)

languages = ['tr', 'es', 'he', 'ar', 'zh', 'ja', 'hi', 'de', 'th', 'ru', 'pl', 'fr', 'pt', 'id', 'bn']

total_items = len(data)
print(f"Toplam {total_items} item çevrilecek...\n")

for idx, item in enumerate(data, 1):
    en_label = item['label']['en']
    en_desc = item['description']['en']
    
    print(f"[{idx}/{total_items}] {en_label[:50]}...")
    
    # Label çevirileri
    for lang_code in languages:
        if not item['label'].get(lang_code):
            try:
                translator = GoogleTranslator(source='en', target=lang_code)
                translated = translator.translate(en_label)
                item['label'][lang_code] = translated
                time.sleep(0.1)  # Rate limiting
            except Exception as e:
                print(f"  ⚠ {lang_code} label hatası: {str(e)[:50]}")
                item['label'][lang_code] = en_label  # Fallback
    
    # Description çevirileri
    for lang_code in languages:
        if not item['description'].get(lang_code):
            try:
                translator = GoogleTranslator(source='en', target=lang_code)
                translated = translator.translate(en_desc)
                item['description'][lang_code] = translated
                time.sleep(0.1)  # Rate limiting
            except Exception as e:
                print(f"  ⚠ {lang_code} desc hatası: {str(e)[:50]}")
                item['description'][lang_code] = en_desc  # Fallback
    
    # Her 5 item'da bir kaydet (güvenlik için)
    if idx % 5 == 0:
        with open('static/data/sexual_preferences.json', 'w', encoding='utf-8') as f:
            json.dump(data, f, ensure_ascii=False, indent=4)
        print(f"  ✓ Kaydedildi ({idx}/{total_items})")
        time.sleep(1)  # API rate limiting

# Final kayıt
with open('static/data/sexual_preferences.json', 'w', encoding='utf-8') as f:
    json.dump(data, f, ensure_ascii=False, indent=4)

print(f"\n✓ Tüm çeviriler tamamlandı! {total_items} item çevrildi.")

