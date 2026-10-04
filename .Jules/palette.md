## 2024-10-03 - Form Label Associations
**Learning:** Missing or mismatched `for` and `id` attributes on form inputs in the app cause screen readers to fail to announce the purpose of the fields.
**Action:** Always ensure `<label for="...">` matches the `<input id="...">` perfectly in the HTML templates.

## 2023-10-25 - Icon-only Buttons and Informative Images Missing A11y Attributes
**Learning:** The application uses icon-only buttons (emoji reactions) and clickable images (QR codes) that lack `aria-label` and `alt` attributes, making them inaccessible to screen readers in this collaborative environment.
**Action:** Always ensure icon-only interactive elements have descriptive `aria-label`s and informative images have clear `alt` text so all users can participate and navigate successfully.
