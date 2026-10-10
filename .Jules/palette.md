## 2024-10-03 - Form Label Associations
**Learning:** Missing or mismatched `for` and `id` attributes on form inputs in the app cause screen readers to fail to announce the purpose of the fields.
**Action:** Always ensure `<label for="...">` matches the `<input id="...">` perfectly in the HTML templates.

## 2023-10-25 - Icon-only Buttons and Informative Images Missing A11y Attributes
**Learning:** The application uses icon-only buttons (emoji reactions) and clickable images (QR codes) that lack `aria-label` and `alt` attributes, making them inaccessible to screen readers in this collaborative environment.
**Action:** Always ensure icon-only interactive elements have descriptive `aria-label`s and informative images have clear `alt` text so all users can participate and navigate successfully.

## 2024-11-20 - Async Action States and Upload UX
**Learning:** When performing asynchronous operations like file uploads, users get confused if the UI immediately dismisses the panel or provides no loading feedback, while screen readers need `aria-busy` to understand the processing state.
**Action:** Always disable submit buttons, provide clear "loading" text, add `aria-busy="true"`, and only dismiss input panels upon confirmed success.

## 2024-12-08 - Async Download Feedback
**Learning:** Users lack confidence when initiating asynchronous downloads if the button remains interactive and provides no immediate feedback.
**Action:** Always add disabled state and `aria-busy` to export/download buttons while fetching data.
## 2024-05-15 - Missing Form Submission Feedback
**Learning:** Found several pages and forms that lack loading states during async tasks (like form submissions in chat), which creates a confusing UX as the user doesn't know if their action is processing.
**Action:** Added a disabled state and 'aria-busy' visual feedback during chat comment submission in `wwwfiles/assets/content.js`, matching the pattern already used for file uploads.
