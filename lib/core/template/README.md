Vangogh Core Template
----

Vangogh uses the `html/template` library in implementing the render engine.

The core render engine assumes a specific layout structure for template definitions.
The root template is `page`. The current template generates valid HTML5 markup.

When adding custom themes using the core render engine, the developer

* Must keep any and all files in the `view` subdirectory where a `content` template
  is defined -- the render engine assumes these file paths as constants. For sake
  of simplicity, changes should be made directly in these files, changing the `content`
  template implementation. The `content` template should not be deleted.
* May replace the `page` template implementation details, as long as the `content`
  template is still called. The minimal `page` template is:
  ```
  {{- define "page" -}}
  {{- template "content" . -}}
  {{- end -}}
  ```
