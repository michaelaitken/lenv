# lenv

**`lenv`** is a cross-platform, CLI-based tool (currently in **proof of concept** stage) I've been developing to help manage local environment variable profiles within a project directory. The idea is to simplify environment management across projects by isolating environment variables locally, so they don't interfere with the global OS-level environment.

> **Note**: `lenv` currently supports **PowerShell on Windows only**. Linux and macOS support is planned but not yet implemented.

## Project Status

`lenv` is still in very early development, and I'm actively building and testing its core features. At this stage, the code is more experimental, and I'm using it to evaluate the concept and test out ideas. As such, the code is not fully optimized or structured for long-term maintainability.

## Features

- **Environment Profiles**: Easily create and manage multiple environment profiles within a project.
- **Local Isolation**: Keeps environment variables project-specific, so they don't impact global system settings.
- **Command Line Interface**: Simple CLI commands for managing profiles, activating, deactivating, and more.

## Planned Commands and Current Support

| Command                      | Description                                                                              | Status                                   |
|------------------------------|------------------------------------------------------------------------------------------|------------------------------------------|
| `lenv help`                  | Display help information                                                                 | Supported, not fully implemented         |
| `lenv init`                  | Initialize a `.lenv` directory for the project                                           | Supported                                |
| `lenv create <profile_name>` | Create a new environment profile                                                         | Supported                                |
| `lenv activate <profile_name>` | Activate the specified profile in the current shell session                           | Supported                                |
| `lenv deactivate`            | Deactivate the current profile and clear environment variables                           | Supported with manual exit instructions  |
| `lenv list`                  | List all available profiles                                                              | Planned                                  |
| `lenv show <profile_name>`   | Display contents of the specified profile                                                | Planned                                  |
| `lenv delete <profile_name>` | Delete a specified profile                                                               | Planned                                  |
| `lenv clone <source_profile> <new_profile>` | Clone an existing profile to a new one                                 | Planned                                  |
| `lenv rename <old_profile_name> <new_profile_name>` | Rename an existing profile                                 | Planned                                  |
| `lenv status`                | Show the currently active profile and status                                             | Planned                                  |
| `lenv import <file_path>`    | Import environment variables from an external file                                       | Planned                                  |
| `lenv diff <profile1> <profile2>` | Compare two profiles for differences                                               | Planned                                  |
| `lenv validate <profile_name>` | Validate the format and integrity of a profile                                         | Planned                                  |

### Known Limitation with `deactivate` Command

Currently, **PowerShell on Windows** does not support session termination via an external Go application. For now, please use the `exit` command manually to deactivate the session:

```powershell
exit
```

I'm looking for alternative solutions to provide a more integrated deactivation process on Windows.

## Installation

Since `lenv` is in the early stages, there isn't a formal installation process yet. To try it out, you can clone the repository and build it manually:

```bash
git clone https://github.com/michaelaitken/lenv
cd lenv
go build -o .\build\lenv.exe .\cmd
```

Then, add the compiled binary to your system’s PATH to make the `lenv` command globally accessible.

## Usage

To get started with `lenv`, open PowerShell in your project directory and initialize the tool with:

```powershell
lenv init
```

Then, create a new profile:

```powershell
lenv create dev
```

To activate the profile:

```powershell
lenv activate dev
```

To deactivate the profile (manually):

```powershell
lenv deactivate # Use 'exit' on Windows. See above
```

Use `lenv help` for more command information.

## Adding Variables to a Profile

To add environment variables to a profile in `lenv`, you can edit the profile's corresponding YAML file. Each profile you create has its own YAML file stored in the `.lenv` directory under your project.

### Steps to Add Variables

1. **Navigate to the Profile File**:
   - After you’ve created a profile (e.g., `dev`), go to the `.lenv` folder in your project directory.
   - Open the profile’s YAML file. For example, if the profile is `dev`, the file will be located at `.lenv/dev.yaml`.

2. **Add Variables in Key-Value Pairs**:
   - Use the YAML format to define your environment variables in key-value pairs. Here’s an example:

   ```yaml
   DATABASE_URL: "postgres://dev_user:dev_pass@localhost/dev_db"
   API_KEY: "dev-12345"
   DEBUG: "true"
   ```

3. **Save the File**:
   - After adding or editing the variables, save the file.

4. **Refresh the Profile**:
   - To update your environment with the new variables, deactivate and then reactivate the profile:
   ```powershell
   lenv deactivate # Use 'exit' on Windows. See above
   lenv activate dev
   ```
   - The session will now reflect any changes made to the variables in `dev.yaml`.

### Example Workflow

If you need to add new variables or update existing ones in the profile, simply edit the YAML file, save it, then run `lenv deactivate` followed by `lenv activate <profile_name>` to load the updated variables.

This approach provides a simple, file-based way to manage environment variables that can easily be version-controlled along with the rest of your project.

## Contributing

If you're interested in contributing, feel free to do so! But please keep in mind that this project is still in its conceptual phase. I'm still testing new ideas, so the structure, features, and overall design may change significantly as development progresses.

### Areas I'm Focusing On

I'm especially interested in contributions that can help with:

- **Linux and macOS compatibility**
- **Optimizing command handling and performance**
- **Finding alternative solutions for deactivation on Windows**

If you'd like to contribute, feel free to fork the repository, make changes, and submit a pull request. I’ll review contributions as development priorities allow.

---

Thanks for checking out `lenv`! If you try it out, let me know how it works for you or if you have suggestions for improvements.
