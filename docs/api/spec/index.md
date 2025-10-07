---
layout: default
title: API Specification
parent: API
nav_order: 1
permalink: /api_spec/
---

<div class="redoc-wide">
  <div id="redoc-container" style="min-height: 85vh;"></div>
</div>

<script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js" defer></script>
<script>
  window.addEventListener('DOMContentLoaded', function () {
    Redoc.init(
      '{{ site.baseurl }}/api_spec/spec/openapi.yaml',
      {
        scrollYOffset: 60,
        hideDownloadButton: false,
        expandResponses: "200,201",
        theme: {
          colors: {
            primary: { main: "#2b6cb0" }
          },
          typography: {
            fontSize: "16px",
            lineHeight: "1.6"
          }
        }
      },
      document.getElementById('redoc-container')
    );
  });
</script>

<style>
.redoc-wide .main-content {
  max-width: 100%;
  padding: 0 1.5rem;
}
</style>
