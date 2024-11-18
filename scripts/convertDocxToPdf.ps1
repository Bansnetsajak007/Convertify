param (
    [string]$inputPath,
    [string]$outputPath
)

# Check if the input file exists
if (-not (Test-Path $inputPath)) {
    Write-Error "Input file does not exist: $inputPath"
    exit 1
}

# Conversion logic using Word COM object
$word = New-Object -ComObject Word.Application
$doc = $word.Documents.Open($inputPath)
$doc.SaveAs([ref]$outputPath, [ref]17)  # 17 represents PDF format
$doc.Close()
$word.Quit()

Write-Host "Conversion completed: $inputPath to $outputPath"
