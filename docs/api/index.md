---
layout: default
title: API Specification
permalink: /api/
---

<div id="redoc-container" style="min-height: 90vh;"></div>

<style>
/* Force Redoc to fill remaining horizontal space next to Just-the-Docs sidebar */
@media (min-width: 1024px) {
  #redoc-container {
    margin-left: 300px;                     /* adjust if your sidebar is narrower/wider */
    width: calc(100% - 300px) !important;
    max-width: none !important;
  }

  /* Loosen JTD’s content clamps so Redoc can stretch fully */
  .container,
  .container-lg,
  .page,
  .page-content,
  .main,
  .main-content {
    max-width: none !important;
    width: 100% !important;
  }

  .main-content {
    padding-left: 1.5rem !important;
    padding-right: 1.5rem !important;
  }
}

.main-content .markdown > *:not(.allow-wide) { max-width: 75ch; }
.allow-wide { max-width: none; }
</style>

<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js" defer></script>
<script>
window.addEventListener('DOMContentLoaded', function () {
  Redoc.init(
    '{{ site.baseurl }}/api/openapi.yaml',
    {
      layout: "stacked",
      scrollYOffset: 60,
      hideDownloadButton: false,
      expandResponses: "200,201",
      theme: {
        colors: { primary: { main: "#2b6cb0" } },
        typography: { fontSize: "16px", lineHeight: "1.6" }
      }
    },
    document.getElementById('redoc-container')
  );
});
</script>
