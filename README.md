# gofast [![Actions Status](https://github.com/johnaoss/gofast/workflows/go/badge.svg)](https://github.com/johnaoss/gofast/actions)


Check your network specs from the menu bar.

# Building

To run this project, run `make` while in the root of this repository.

This was tested to work on macOS Catalina Beta (19A546d).

# Roadmap

Actually implement the project. Deciding on either using `fast.com` or the internal macOS network speeds, but that requires me to write some Objective-C which I'd rather not.

# License

Please check [LICENSE.md](LICENSE.md) for information about the license.

# Acknowledgements

This project is mainly based off of [caseymrm/menuet](github.com/caseymrm/menuet), a macOS menu bar library.

The app's icon is "Gauge" by Icons Bazaar from [the Noun Project](https://thenounproject.com/search/?q=gauge&i=624881).

Check your network specs from the menu bar.

# Building

To run this project, run `make` while in the root of this repository.

This was tested to work on macOS Catalina Beta (19A546d).

# Roadmap

The base level of the project is currently implemented, however I'd like to rewrite the library I use to obtain the actual `fast.com` speed from. 

Additionally, I plan on providing a history that can be looked at over time, as well as potentially scheduled speedtests depending on current network load. This would require me to hook into the system-level network speeds, and as such I'd need to learn a bit of Objective-C first.

The final big TODO is to actually make the app look pretty, though I'm unsure of how to do that without using native Cocoa resources, of which there isn't a large wrapper for. 

# License

Please check [LICENSE.md](LICENSE.md) for information about the license.

# Acknowledgements

This project is mainly based off of [caseymrm/menuet](github.com/caseymrm/menuet), a macOS menu bar library.

The app's icon is "Gauge" by Icons Bazaar from [the Noun Project](https://thenounproject.com/search/?q=gauge&i=624881).
