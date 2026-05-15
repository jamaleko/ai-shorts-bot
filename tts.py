import edge_tts
import asyncio
import sys

TEXT = sys.argv[1]

async def main():

    communicate = edge_tts.Communicate(
        TEXT,
        "en-US-GuyNeural"
    )

    await communicate.save("voice.mp3")

asyncio.run(main())
