---
layout: default
title: API Specification
permalink: /api/
---

<div id="redoc-container" style="min-height: 85vh;"></div>

<style>
/* Force Redoc page full width */
@media (min-width: 1024px) {
  /* Blow away every JTD wrapper width */
  .container,
  .container-lg,
  .page,
  .page-content,
  .main,
  .main-content,
  .content,
  .wrap {
    max-width: none !important;
    width: 100% !important;
  }

  .main-content {
    padding-left: 1.5rem !important;
    padding-right: 1.5rem !important;
  }

  /* Stretch Redoc fully */
  #redoc-container {
    max-width: none !important;
    width: 100% !important;
    min-height: 85vh;
  }
}
</style>

<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js" defer></script>
<script>
window.addEventListener('DOMContentLoaded', function () {
  Redoc.init(
    '{{ site.baseurl }}/api/openapi.yaml',
    {
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
