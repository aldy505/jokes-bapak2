<script lang="ts">
  import { _ } from 'svelte-i18n'
  import { goto } from '$app/navigation'
  import env from '$lib/env';
  import Codeblock from '../components/codeblock.svelte';
</script>

<svelte:head>
  <title>{$_('meta.title')} - {$_('meta.tagline')}</title>
</svelte:head>

<section>
  <div class="flex flex-col md:flex-row items-center py-8">
    <div class="flex-1">
      <h1 class="text-6xl font-bold py-2">{$_('meta.tagline')}</h1>
      <p class="text-base py-4 w-2/3">{$_('meta.explanation')}</p>
    </div>
    <div class="flex-1 px-6">
      <div class="max-w-xs mx-auto">
        <img src={env.API_ENDPOINT + `/today`} alt="Sample joke" class="py-6 shadow-2xl">
      </div>
      <Codeblock class="px-4 md:px-0 mx-auto">
        $ curl -XGET 'https://jokesbapak2.herokuapp.com/v1/'
      </Codeblock>
      <p class="text-sm text-center py-4 opacity-70 hover:opacity-100 transition duration-300 ease-in-out">
        {$_('home.more.1')} <span on:click={() => goto('/guide')}>{$_('navigation.guide')}</span> {$_('home.more.2')} <span on:click={() => goto('/api')}>{$_('navigation.api')}</span>
      </p>
    </div>
  </div>
</section>

<style>
  span {
    @apply 'hover:underline' cursor-pointer;
  }
</style>
