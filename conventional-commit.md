To create a PR please follow conventional commit based on the specification from conventionalcommits.org, you need to follow these guidelines:

The commit message format should follow this structure: <type>(<scope>): <subject>. Angle brackets indicate optional fields.

<type>: Describes the purpose of the commit. It can be one of the following:
feat: A new feature
fix: A bug fix
docs: Documentation changes
style: Formatting, missing semi-colons, etc.
refactor: Code changes that neither fix a bug nor add a feature
test: Adding missing tests or correcting existing tests
chore: Other changes that don't modify source or test files
<scope>: Optional field that specifies the scope of the change. It can be a component, file, or module name.
<subject>: A concise description of the change in the present tense. It should not exceed 50 characters.
Example commit message: feat(login): Add forgot password feature

Optionally, you can provide a longer commit message in the body section. It should provide more details about the change. This section starts after a blank line following the subject. It should be wrapped within 72 characters per line.

Example commit message with body:

feat(login): Add forgot password feature

- Add forgot password form
- Implement password reset logic
- Update UI for password recovery process
- Write unit tests for new functionality
Copy Code
Additionally, you can include footers to refer to any issues or provide additional information.

Example commit message with footer:

feat(login): Add forgot password feature

- Add forgot password form
- Implement password reset logic
- Update UI for password recovery process
- Write unit tests for new functionality

Closes #123
Copy Code
Remember to replace <type>, <scope>, and <subject> with appropriate values based on the changes you are making. Following this format will help maintain a clean and organized commit history.