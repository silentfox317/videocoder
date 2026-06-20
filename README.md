<div align="center">

# 🎬 VideoCoder

**AI video editor in terminal chat**

Just describe what you want — VideoCoder edits your video using MCP-video tools (81 tools) and shell commands.

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Based on](https://img.shields.io/badge/Based%20on-OpenCode-000?logo=openai)](https://github.com/opencode-ai/opencode)

---

https://github.com/user-attachments/assets/placeholder

*Talk to AI. Edit video. Done.*

</div>

## ✨ Features

- **🎯 Chat-based editing** — describe what you want in plain language
- **🔧 81 MCP-video tools** — trim, merge, subtitles, audio, effects, stabilize, transcribe
- **🖥️ Shell access** — custom ffmpeg scripts when you need them
- **🧠 Any LLM** — works with Mistral, OpenAI, Anthropic, local models
- **📱 Social media presets** — YouTube, TikTok, Instagram, Reels

## 🚀 Quick Start

```bash
# Download
git clone https://github.com/silentfox317/videocoder.git
cd videocoder

# Build
go build -o videocoder .

# Configure
cp .videocoder.example.json .videocoder.json
# Edit .videocoder.json with your LLM provider settings

# Run
./videocoder
```

## 📖 Usage

Just type what you want to do with your video:

```
> Trim the first 10 seconds from my video.mp4
> Merge all clips from the ./raw folder
> Add subtitles in Russian to this video
> Extract the audio track as MP3
> Split this long recording into chapters
> Transcribe the speech in this video
```

## 🛠️ MCP Video Tools

VideoCoder comes with **81 MCP-video tools** including:

| Category | Tools |
|----------|-------|
| 📐 **Basic** | trim, merge, split, concatenate |
| 📝 **Text** | subtitles, captions, text overlay |
| 🎵 **Audio** | extract, replace, mix, volume |
| ✨ **Effects** | filters, color grade, transitions |
| 📐 **Resize** | crop, scale, aspect ratio |
| 🔄 **Format** | transcode, compress, repurpose |
| 🎯 **Analysis** | inspect, detect scenes, transcribe |
| 🧊 **Stabilize** | stabilization, rolling shutter fix |
| 🌐 **Social** | TikTok preset, Reels, YouTube Shorts |

## ⚙️ Configuration

See `.videocoder.example.json` for all options.

### LLM Providers

```json
{
  "providers": {
    "openai-compatible": {
      "model": "gpt-4o",
      "baseUrl": "http://localhost:5000/v1",
      "apiKey": "your-key"
    }
  }
}
```

### MCP Server

VideoCoder uses `@kyanite/mcp-video` by default. Install it:

```bash
npx @kyanite/mcp-video
```

## 🏗️ Architecture

VideoCoder is a fork of [OpenCode](https://github.com/opencode-ai/opencode) by OpenAI, repurposed for video editing.

- **Chat TUI** — Bubble Tea terminal UI
- **LLM Agent** — tool-calling loop
- **MCP Client** — connects to mcp-video server
- **Shell** — ffmpeg access

## 📄 License

MIT License — see [LICENSE](LICENSE)

---

*Made for video creators who'd rather talk than click*

</div>
