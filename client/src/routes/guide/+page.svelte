<script lang="ts">
  import Codeblock from '../../components/codeblock.svelte';
  import { _ } from 'svelte-i18n';
  import env from '$lib/env';
  // This page is meant to guide people on how to use the API.
</script>

<svelte:head>
  <title>{$_('navigation.guide')} - {$_('meta.title')}</title>
  <meta content={$_('navigation.guide') + '-' + $_('meta.title')} name="title" />
  <meta content={$_('navigation.guide') + '-' + $_('meta.title')} name="twitter:title" />
  <meta content={$_('navigation.guide') + '-' + $_('meta.title')} property="og:title" />
  <link href="https://jokesbapak2.reinaldyrafli.com/guide" rel="canonical" />
  <meta content="Largest collection of Indonesian dad jokes as a consumable API" name="description" />
  <meta content="Largest collection of Indonesian dad jokes as a consumable API" name="twitter:description" />
  <meta content="Largest collection of Indonesian dad jokes as a consumable API" property="og:description" />
</svelte:head>

<section class="guide_page">
  <h1>{$_('navigation.guide')}</h1>
  <p>
    {$_('guide.introduction.1')}
    {$_('guide.introduction.2')}
  </p>
</section>

<section class="guide_page">
  <h2>{$_('guide.direct.1')} <code>&lt;img&gt;</code> {$_('guide.direct.2')}</h2>
  <Codeblock>&lt;img src="{env.BROWSER_API_ENDPOINT}/" /&gt;</Codeblock>
</section>

<section class="guide_page">
  <h2>{$_('guide.fetch.1')}</h2>
  <p>{$_('guide.fetch.2')}</p>
  <Codeblock>
    const response = await fetch(&quot;{env.BROWSER_API_ENDPOINT}/&quot;);<br />
    <br />
    if (!response.ok) &#123;<br />
    &nbsp;&nbsp;// {$_('guide.fetch.3')}<br />
    &#125;<br />
    <br />
    const blob = await response.blob();<br />
    <br />
    const objectURL = URL.createObjectURL(blob);<br />
    <br />
    &lt;img src=&quot;&#123; objectURL &#125;&quot; /&gt;
  </Codeblock>
</section>

<style>
  p {
    @apply text-base;
    @apply py-2;
    @apply lg\:w-2\/3;
  }

  h1 {
    @apply text-4xl;
    @apply font-bold;
    @apply py-2;
  }

  h2 {
    @apply text-2xl;
    @apply font-bold;
    @apply py-2;
  }

  a {
    @apply hover\:underline;
    @apply dark\:text-dodger-200;
    @apply text-dodger-700;
  }

  section {
    @apply pt-6;
  }
</style>
