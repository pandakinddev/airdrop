$loc=(Get-Location).Path -replace "\\","/"

$tokenRoot=(Join-Path -Path $loc -ChildPath coin) -replace "\\","/"
$contracts = (Join-Path -Path $tokenRoot -ChildPath contracts) -replace "\\","/"

Get-ChildItem -Path $contracts -File -Recurse | `
    ForEach-Object { 
        $content=(Get-Content -path $_.FullName -Raw) -replace '"@openzeppelin/',"`"$tokenRoot/node_modules/@openzeppelin/"
        $content | Out-File -FilePath $_.FullName -Encoding utf8
    }

# solc --abi $contracts/TimedAirdrop.sol `
#     $tokenRoot/node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol `
#     -o abi --combined-json abi

solc --abi $contracts/test/SimpleAirdrop.sol `
    $contracts/Airdrop.sol `
    $tokenRoot/node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol `
    -o abi --combined-json abi

abigen.exe --combined-json `
    $loc/abi/combined.json --pkg airdrop `
    --out $loc/pkg/airdrop/airdrop.go

Remove-Item -Path $loc/abi -Recurse -Force