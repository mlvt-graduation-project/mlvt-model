# stt.py

import whisper
import os
import argparse
import sys

def process_stt(input_file, output_file):
    try:
        # Load the Whisper model
        model = whisper.load_model("base")

        if not os.path.isfile(input_file):
            print(f"Input file '{input_file}' does not exist.", file=sys.stderr)
            sys.exit(1)

        result = model.transcribe(input_file)

        with open(output_file, "w") as file:
            file.write(result["text"])

        print(f"Transcription for '{input_file}' saved to '{output_file}'")

    except Exception as e:
        print(f"An error occurred during STT processing: {str(e)}", file=sys.stderr)
        sys.exit(1)

def main():
    parser = argparse.ArgumentParser(description="STT Processing Script")
    parser.add_argument('--input_file', type=str, required=True, help="Path to the input video file.")
    parser.add_argument('--output_file', type=str, required=True, help="Path to the output transcription text file.")

    args = parser.parse_args()

    input_file = args.input_file
    output_file = args.output_file

    process_stt(input_file, output_file)

if __name__ == "__main__":
    main()
