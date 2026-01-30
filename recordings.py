import os
import shutil
import speech_recognition as sr
from langdetect import detect, LangDetectException

# Configurable paths
SOURCE_FOLDER = r'C:\path\to\your\recordings'
DEST_FOLDER = r'C:\path\to\hindienglish_recording'

os.makedirs(DEST_FOLDER, exist_ok=True)
recognizer = sr.Recognizer()

def is_english_or_hindi(text):
    try:
        lang = detect(text)
        return lang in ['en', 'hi']
    except LangDetectException:
        return False

for filename in os.listdir(SOURCE_FOLDER):
    if filename.lower().endswith(('.wav', '.mp3', '.flac', '.m4a')):
        file_path = os.path.join(SOURCE_FOLDER, filename)
        try:
            with sr.AudioFile(file_path) as source:
                audio = recognizer.record(source)
                text = recognizer.recognize_google(audio, language='en-IN')
                if is_english_or_hindi(text):
                    shutil.move(file_path, os.path.join(DEST_FOLDER, filename))
                    print(f"Moved: {filename}")
        except Exception as e:
            print(f"Could not process {filename}: {e}")
