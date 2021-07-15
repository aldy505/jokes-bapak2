<script lang="ts">
  // This page is meant to explain available API endpoints.
  import { _ } from 'svelte-i18n'
  import env from '$lib/env'
  import Codeblock from '../components/codeblock.svelte'
</script>

<svelte:head>
  <title>{$_('navigation.api')} - {$_('meta.title')}</title>
</svelte:head>

<section>
  <h1>{$_('api.get.title')}</h1>
  <h2>{$_('api.get.random.title')}</h2>
  <p>{$_('api.get.random.body')}</p>
  <Codeblock>
    GET {env.API_ENDPOINT}/
  </Codeblock>
  <h2>{$_('api.get.today.title')}</h2>
  <p>{$_('api.get.today.body')}</p>
  <Codeblock>
    GET {env.API_ENDPOINT}/today
  </Codeblock>
  <h2>{$_('api.get.id.title')}</h2>
  <p>{$_('api.get.id.body')}</p>
  <Codeblock>
    GET {env.API_ENDPOINT}/&lcub;id&rcub;
  </Codeblock>
</section>

<section>
  <p class="pt-8">{$_('api.more')}</p>
</section>

<style>
  h1 {
    @apply text-4xl font-bold py-4;
  }
  h2 {
    @apply text-2xl font-bold pt-6 pb-1;
  }
  p {
    @apply text-base opacity-80 py-2;
  }
</style>