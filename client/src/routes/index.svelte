<script context="module">
  export const prerender = true;
</script>

<script lang="ts">
  import { _ } from 'svelte-i18n';
  import { goto } from '$app/navigation';
  import env from '$lib/env';
  import Codeblock from '../components/codeblock.svelte';
</script>

<svelte:head>
  <title>{$_('meta.title')} - {$_('meta.tagline')}</title>
  <meta name="twitter:title" content={$_('meta.title') + '-' + $_('meta.tagline')} />
  <meta property="og:title" content={$_('meta.title') + '-' + $_('meta.tagline')} />
  <link rel="canonical" href="https://jokesbapak2.pages.dev/" />
  <meta name="twitter:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
  <meta property="og:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
</svelte:head>

<section>
  <div class="flex flex-col lg:flex-row items-center py-8">
    <div class="flex-1">
      <h1 class="text-5xl md:text-6xl font-bold py-2">{$_('meta.tagline')}</h1>
      <p class="text-base py-4 md:w-2/3">{$_('meta.explanation')}</p>
    </div>
    <div class="flex-1 md:px-6 w-full">
      <div class="max-w-xs mx-auto">
        <img src={env.API_ENDPOINT + `/today`} alt="Sample joke" class="py-6 shadow-2xl" />
      </div>
      <Codeblock>$ curl -XGET 'https://jokesbapak2.herokuapp.com/v1/'</Codeblock>
      <p class="text-sm text-center py-4 opacity-70 hover:opacity-100 transition duration-300 ease-in-out">
        {$_('home.more.1')} <span on:click={() => goto('/guide')}>{$_('navigation.guide')}</span>
        {$_('home.more.2')} <span on:click={() => goto('/api')}>{$_('navigation.api')}</span>
      </p>
    </div>
  </div>
</section>

<style>
  span {
    @apply 'hover:underline' cursor-pointer;
  }
</style>
