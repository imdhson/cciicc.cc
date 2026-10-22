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

## 2024-12-10 - Visually Hidden Interactive Elements Focusability
**Learning:** Using `max-height: 0` and `opacity: 0` on an interactive element panel (like an upload area) visually hides it but leaves its children focusable by keyboards and screen readers, leading to confusing navigation in the app.
**Action:** Always add `visibility: hidden` (and `visibility: visible` when shown) or `display: none` to panels that are visually hidden, so their contents are correctly removed from the accessibility and focus tree.

## 2024-12-15 - ARIA attributes on interactive toggles
**Learning:** Toggling visual state (like expanding an image) without updating ARIA attributes leaves screen reader users unaware of the new state and the available action.
**Action:** Always dynamically update `aria-expanded` and `aria-label` (e.g., changing 'view larger' to 'view smaller') via JavaScript when implementing togglable interactive elements.

## 2025-01-21 - Pagination Boundaries UX
**Learning:** Users lack visual cues when they reach the beginning or end of a PDF document because the navigation buttons remain fully active, and screen readers are not notified of page changes because the page number text lacks an `aria-live` region.
**Action:** Always disable pagination buttons (`prev`/`next`) at the boundaries of the content, style them appropriately to reflect the disabled state, and add `aria-live="polite"` to dynamically changing text like page numbers.

## 2024-05-13 - A11y updates for dynamic text
**Learning:** Screen readers won't automatically announce new toasts or dynamically added DOM elements (like chat bubbles) unless they're properly wrapped in an ARIA live region (`aria-live="polite"` or `aria-live="assertive"`).
**Action:** When creating container elements meant for dynamic user-facing updates (like `<div id="popup">` or `<div id="message-container">`), always add `aria-live="polite"` and `aria-atomic="true"`.
