# **Convertify**

Convertify is a versatile, interactive command-line tool written in Go for converting and compressing media and documents. It supports seamless document transformations and image operations like compression and format conversion.

---

## **Features**
- **Document Conversion**:
  - Convert DOCX to PDF.
  - Convert PDF to DOCX.
- **Image Compression**:
  - Compress JPG images with adjustable quality.
- **Image Format Conversion**:
  - Convert JPG to PNG.
  - Convert PNG to JPG.

---

## **Installation**
1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd Convertify
   ```

2. **Build the application**:
   ```bash
   go build -o convertify
   ```

3. **Run the tool**:
   ```bash
   ./convertify help
   ```

## **Usage**

### **1. Document Conversion**
* Convert DOCX to PDF:
  ```bash
  convertify docsconvert <input.docx> <output.pdf>
  ```
* Convert PDF to DOCX:
  ```bash
  convertify docsconvert <input.pdf> <output.docx>
  ```

### **2. Image Compression**
* Compress a JPG image:
  ```bash
  convertify imgCompress <input.jpg> <output.jpg>
  ```

### **3. Image Format Conversion**
* Convert JPG to PNG:
  ```bash
  convertify imgConvert <input.jpg> <output.png>
  ```
* Convert PNG to JPG:
  ```bash
  convertify imgConvert <input.png> <output.jpg>
  ```

## **Notes**
1. Replace `<input>` and `<output>` with your actual file paths.
2. Ensure read/write permissions for the input and output directories.
3. Powershell scripts for document conversion must be configured and accessible in the project directory.

## **Future Enhancements** 🚀

### **1. Video Processing**
- [ ] Video compression with configurable quality
- [ ] Format conversion (MP4, AVI, MKV, etc.)
- [ ] Video trimming and merging
- [ ] Extract audio from video
- [ ] Batch video processing

### **2. Audio Processing**
- [ ] Audio format conversion (MP3, WAV, AAC, etc.)
- [ ] Audio compression
- [ ] Basic audio editing (trim, merge)
- [ ] Metadata management

### **3. Document Enhancements**
- [ ] Support for more document formats (ODT, RTF, etc.)
- [ ] Batch document conversion
- [ ] Document merging
- [ ] PDF compression

### **4. Image Processing Extensions**
- [ ] Support for more image formats (WEBP, GIF, TIFF)
- [ ] Batch image processing
- [ ] Basic image editing (resize, crop, rotate)
- [ ] Image optimization for web
- [ ] Metadata preservation options

### **5. Cloud Integration**
- [ ] Direct upload/download from cloud storage
- [ ] Integration with popular cloud services (Google Drive, Dropbox)
- [ ] Cloud-based processing options

### **6. Advanced Features**
- [ ] Custom conversion presets
- [ ] Automated quality optimization
- [ ] File integrity verification
- [ ] Compression statistics and reports
- [ ] Command completion for shells

## **Contributing**
Contributions are welcome! Feel free to submit issues or pull requests.

## **Acknowledgments**
Thank you for choosing Convertify. Your feedback helps us improve!