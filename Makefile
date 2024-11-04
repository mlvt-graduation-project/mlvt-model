# Makefile mẹ

.PHONY: tts ttt tts ls

# Định nghĩa biến
PYTHON=python3

tts:
	cd ./scripts/tts && $(PYTHON) tts.py $(INPUT_FILE) $(OUTPUT_FILE)

ttt:
	cd ./scripts/ttt && $(PYTHON) ttt.py $(INPUT_FILE) $(OUTPUT_FILE) $(SOURCE_LANGUAGE) $(TARGET_LANGUAGE)

stt:
	cd ./scripts/stt && $(PYTHON) stt.py $(INPUT_FILE) $(OUTPUT_FILE)

# ls:
# 	cd ./scripts/ls && $(PYTHON) ls.py $(INPUT_FILE) $(OUTPUT_FILE) $(SOURCE_LANGUAGE) $(TARGET_LANGUAGE)
