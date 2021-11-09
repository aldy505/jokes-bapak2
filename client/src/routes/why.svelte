<script lang="ts">
  import { _ } from 'svelte-i18n';
</script>

<svelte:head>
  <title>{$_('navigation.why')} - {$_('meta.title')}</title>
  <meta name="title" content={$_('navigation.why') + '-' + $_('meta.title')} />
  <meta name="twitter:title" content={$_('navigation.why') + '-' + $_('meta.title')} />
  <meta property="og:title" content={$_('navigation.why') + '-' + $_('meta.title')} />
  <link rel="canonical" href="https://jokesbapak2.pages.dev/why" />
  <meta name="description" content="Largest collection of Indonesian dad jokes as a consumable API" />
  <meta name="twitter:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
  <meta property="og:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
</svelte:head>

<section class="why_page">
  <h1 id="why-does-this-project-exists">{$_('why.exists.title')}</h1>
  <p>{$_('why.exists.body.1')} {$_('why.exists.body.2')} {$_('why.exists.body.3')} {$_('why.exists.body.4')}</p>
  <p>{$_('why.exists.body.5')} {$_('why.exists.body.6')}</p>
</section>
<section class="why_page">
  <h1 id="can-i-submit-my-dad-joke">{$_('why.submit.title')}</h1>
  <p>
    {$_('why.submit.body.1')}
    {$_('why.submit.body.2')}
    <a href="mailto:aldy505@tutanota.com">{$_('why.submit.body.3')}</a>
    {$_('why.submit.body.4')}
    {$_('why.submit.body.5')}
    {$_('why.submit.body.6')}
  </p>
</section>
<section class="why_page">
  <h1 id="can-i-contribute">{$_('why.contribute.title')}</h1>
  <p>
    {$_('why.contribute.body.1')}
    {$_('why.contribute.body.2')}
    <a href="https://www.github.com/aldy505/jokes-bapak2">{$_('why.contribute.body.3')}</a>
    {$_('why.contribute.body.4')}
  </p>
</section>
<section class="why_page">
  <h1 id="other-inquiries">{$_('why.inquiries.title')}</h1>
  <p>
    {$_('why.inquiries.body.1')}
    <a href="mailto:aldy505@tutanota.com">{$_('why.inquiries.body.2')}</a>, {$_('why.inquiries.body.3')}
  </p>
</section>

<style>
  p {
    @apply text-base;
    @apply py-2;
    @apply lg\:w-2\/3;
  }

  h1 {
    @apply text-3xl;
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
