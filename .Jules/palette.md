## 2024-10-03 - Form Label Associations
**Learning:** Missing or mismatched `for` and `id` attributes on form inputs in the app cause screen readers to fail to announce the purpose of the fields.
**Action:** Always ensure `<label for="...">` matches the `<input id="...">` perfectly in the HTML templates.

## 2023-10-25 - Icon-only Buttons and Informative Images Missing A11y Attributes
**Learning:** The application uses icon-only buttons (emoji reactions) and clickable images (QR codes) that lack `aria-label` and `alt` attributes, making them inaccessible to screen readers in this collaborative environment.
**Action:** Always ensure icon-only interactive elements have descriptive `aria-label`s and informative images have clear `alt` text so all users can participate and navigate successfully.

## 2024-11-20 - Async Action States and Upload UX
**Learning:** When performing asynchronous operations like file uploads, users get confused if the UI immediately dismisses the panel or provides no loading feedback, while screen readers need `aria-busy` to understand the processing state.
**Action:** Always disable submit buttons, provide clear "loading" text, add `aria-busy="true"`, and only dismiss input panels upon confirmed success.
