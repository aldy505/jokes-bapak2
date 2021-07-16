<script lang="ts">
  import Codeblock from '../components/codeblock.svelte';
  import { _ } from 'svelte-i18n';
  // This page is meant to guide people on how to use the API.
</script>

<svelte:head>
  <title>{$_('navigation.guide')} - {$_('meta.title')}</title>
  <meta name="twitter:title" content={$_('navigation.guide') + '-' + $_('meta.title')} />
  <meta property="og:title" content={$_('navigation.guide') + '-' + $_('meta.title')} />
  <link rel="canonical" href="https://jokesbapak2.pages.dev/guide" />
  <meta name="description" content="Largest collection of Indonesian dad jokes as a consumable API" />
  <meta name="twitter:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
  <meta property="og:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
</svelte:head>

<section>
  <h1>{$_('navigation.guide')}</h1>
  <p>
    {$_('guide.introduction.1')}
    {$_('guide.introduction.2')}
  </p>
</section>

<section>
  <h2>{$_('guide.direct.1')} <code>&lt;img&gt;</code> {$_('guide.direct.2')}</h2>
  <Codeblock>&lt;img src="https://jokesbapak2.herokuapp.com/v1/" /&gt;</Codeblock>
</section>

<section>
  <h2>{$_('guide.fetch.1')}</h2>
  <p>{$_('guide.fetch.2')}</p>
  <Codeblock>
    fetch("https://jokesbapak2.herokuapp.com/v1/")<br />
    &nbsp;&nbsp;.then(response => response.text())<br />
    &nbsp;&nbsp;.then(data => console.log(data))<br />
    <br />
    // {$_('guide.fetch.3')}<br />
    <br />
    &lt;img src="data:image/jpeg;base64, &lcub; data &rcub;" /&gt;
  </Codeblock>
</section>

<style>
  p {
    @apply text-base py-2 'lg:w-2/3';
  }
  h1 {
    @apply text-4xl font-bold py-2;
  }
  h2 {
    @apply text-2xl font-bold py-2;
  }
  a {
    /* This would probably be an error if you installed WindiCSS extension on VSCode */
    @apply 'hover:underline' 'dark:text-dodger-200' text-dodger-700;
  }
  section {
    @apply pt-6;
  }
</style>
