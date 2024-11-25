param (
    [string]$inputPath,
    [string]$outputPath
)

# Check if the input file exists
if (-not (Test-Path $inputPath)) {
    Write-Error "Input file does not exist: $inputPath"
    exit 1
}

# Add COM object for Microsoft Word
Add-Type -AssemblyName "Microsoft.Office.Interop.Word"

# Create a new Word application instance and set it to invisible
$word = New-Object -ComObject Word.Application
$word.Visible = $false

try {
    # Open the PDF document
    $document = $word.Documents.Open($inputPath)
    
    # Save the document as DOCX
    $document.SaveAs([ref]$outputPath, [ref]16)  # 16 represents the default DOCX format
    
    # Close the document
    $document.Close()
    
    Write-Host "Conversion completed: $inputPath to $outputPath"
} catch {
    Write-Error "An error occurred during conversion: $_"
} finally {
    # Quit the Word application
    $word.Quit()
}
