# Password Generator

A flexible password generator written in Go that creates secure passwords with customizable complexity levels and character types.

## Features

- **Three hardness levels**: Easy, Medium, Hard
- **Customizable character types**: Include/exclude digits and special symbols
- **Random generation**: Uses cryptographically secure random number generator
- **Configurable digit count**: Specify how many digits to include when enabled

## Installation

```bash
git clone https://github.com/sellnft/password-generator.git
cd 
go build -o password-generator
```

## Usage

# Run without flags (defaults: digits=true, symbols=true, hardness=medium)
```bash
./password-generator
```

# Custom hardness level
```bash
./password-generator -hardness=easy
./password-generator -hardness=hard
```

# Disable digits or symbols
```bash
./password-generator -digits=false
./password-generator -symbols=false
```

# Disable both digits and symbols
```bash
./password-generator -digits=false -symbols=false
```

| Hardness | Password Length | Special Symbols Count |
|----------|-----------------|-----------------------|
| easy     | 7               | 2                     |
| medium   | 15              | 4                     |
| hard     | 23              | 6                     |

## Examples
# Generate a medium password with digits and symbols (default)
```bash
./password-generator
```
# Output: aB3cD5eFgH!jK7lM9nPqR@t

# Generate an easy password without symbols
```bash
./password-generator -hardness=easy -symbols=false
```
# Output: Gt7nBq2

# Generate a hard password with only letters
```bash
./password-generator -hardness=hard -digits=false -symbols=false
```
# Output: KsLdQwErTyZuIoPaSdFgHj
