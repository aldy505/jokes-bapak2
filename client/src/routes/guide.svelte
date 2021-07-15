<script lang="ts">
  import { _ } from 'svelte-i18n';
  // This page is meant to guide people on how to use the API.
</script>

<svelte:head>
  <title>{$_('navigation.guide')} - {$_('meta.title')}</title>
  <meta name="twitter:title" content={$_('navigation.guide') + '-' + $_('meta.title')} />
  <meta property="og:title" content={$_('navigation.guide') + '-' + $_('meta.title')} />
  <link rel="canonical" href="https://jokesbapak2.pages.dev/guide" />
  <meta name="twitter:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
  <meta property="og:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
</svelte:head>

<section>
  <p class="text-lg py-4">
    Sorry for the inconvinience. For now, please refer to the <a class="underline" href="/api">API</a> first. I'll work on
    this page later on.
  </p>
</section>
