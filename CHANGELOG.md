
# Change Log
All changes to vilmos language and vilmos interpreter will be documented in this file.
 
The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [2.0.1] - 2021-11-07
Beautified code  
 
### Added
### Changed
### Fixed
 
- Fixe import bugs cause by major versione update

## [2.0.0] - 2021-11-07
Scaled images support, file management and strings support

### Added

- Support for scaled images
- Support for strings type with delimiter feature. Read [language specifications](https://github.com/Vinetwigs/vilmos/blob/main/LANGUAGE.md) for more informations about
- pixel_size flag to specify each pixel dimension (to use when each instruction is bigger than 1 px)
- FILE_OPEN instruction to open files
- FILE_CLOSE instruction to close an open file
- Added FILE_OPEN and FILE_CLOSE instructions to configs.ini
 
### Changed

- If a file is opened, INPUT_ASCII puts into the stack all its content
- If a file is opened, INPUT_INT puts into the stack all its content
- If a file is opened, OUTPUT_ASCII writes stack content into the file as chars
- If a file is opened, OUTPUT_INT writes stack content into the file as int
- OUTPUT_INT hex value is now set to #000001
- Updated README, LANGUAGE.md, CHANGELOG.md

Read [language specifications](https://github.com/Vinetwigs/vilmos/blob/main/LANGUAGE.md) for more detailed informations

### Fixed
- General code refactor
- Fixed LANGUAGE.md color preview for all the instructions

## [1.2.0] - 2021-11-06
New INPUT_ASCII behaviour

### Added

- INPUT_ASCII operation now can push into the stack an entire string and not an unique char
 
### Changed
### Fixed
 
## [1.1.0] - 2021-11-03
Beautified code  
 
### Added
### Changed
### Fixed
 
- Code refactor - by [Dar9586](https://github.com/Dar9586)
 
## [1.0.0] - 2021-11-03
 
- Initial Release
