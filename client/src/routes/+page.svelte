<script lang="ts">
  import { _ } from 'svelte-i18n';
  import { goto } from '$app/navigation';
  import env from '$lib/env';
  import Codeblock from '../components/codeblock.svelte';

  /** @type {import('./$types').PageData} */
  export let data;

  let total = data.total;
</script>

<svelte:head>
  <title>{$_('meta.title')} - {$_('meta.tagline')}</title>
  <meta content={$_('meta.title') + '-' + $_('meta.tagline')} name="title" />
  <meta content={$_('meta.title') + '-' + $_('meta.tagline')} name="twitter:title" />
  <meta content={$_('meta.title') + '-' + $_('meta.tagline')} property="og:title" />
  <link href="https://jokesbapak2.reinaldyrafli.com/" rel="canonical" />
  <meta content="Largest collection of Indonesian dad jokes as a consumable API" name="description" />
  <meta content="Largest collection of Indonesian dad jokes as a consumable API" name="twitter:description" />
  <meta content="Largest collection of Indonesian dad jokes as a consumable API" property="og:description" />
</svelte:head>

<section>
  <div class="flex flex-col lg:flex-row items-center py-8">
    <div class="flex-1">
      <h1 class="text-4xl sm:text-5xl md:text-6xl font-bold py-2">{$_('meta.tagline-total', { values: { total } })}</h1>
      <p class="text-base py-4 md:w-2/3">{$_('meta.explanation')}</p>
    </div>
    <div class="flex-1 md:px-6 w-full">
      <div class="max-w-xs mx-auto">
        <img alt="Sample joke" class="py-6 shadow-2xl" src={env.BROWSER_API_ENDPOINT + `/today`} />
      </div>
      <Codeblock>$ curl -XGET 'https://jokesbapak2.reinaldyrafli.com/api/'</Codeblock>
      <p class="text-sm text-center py-4 opacity-70 hover:opacity-100 transition duration-300 ease-in-out">
        {$_('home.more.1')}
        <span class="hover:underline cursor-pointer" on:click={() => goto('/guide')}>{$_('navigation.guide')}</span>
        {$_('home.more.2')}
        <span class="hover:underline cursor-pointer" on:click={() => goto('/api')}>{$_('navigation.api')}</span>
      </p>
    </div>
  </div>
</section>
