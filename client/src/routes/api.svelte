<script lang="ts">
  // This page is meant to explain available API endpoints.
  import { onMount } from 'svelte';
  import { _ } from 'svelte-i18n';
  import env from '$lib/env';
  import { $fetch as omf } from 'ohmyfetch';
  import Codeblock from '../components/codeblock.svelte';
  import Notice from '../components/notice.svelte';

  interface TotalResponse {
    message: string;
  }

  let total;

  onMount(async () => {
    const totalJokes = async (): Promise<string> => {
      const response = await omf<TotalResponse>(`${env.API_ENDPOINT}/total`);
      return response.message;
    };

    total = await totalJokes();
  });
</script>

<svelte:head>
  <title>{$_('navigation.api')} - {$_('meta.title')}</title>
  <meta name="title" content={$_('navigation.api') + '-' + $_('meta.title')} />
  <meta name="twitter:title" content={$_('navigation.api') + '-' + $_('meta.title')} />
  <meta property="og:title" content={$_('navigation.api') + '-' + $_('meta.title')} />
  <link rel="canonical" href="https://jokesbapak2.pages.dev/api" />
  <meta name="description" content="Largest collection of Indonesian dad jokes as a consumable API" />
  <meta name="twitter:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
  <meta property="og:description" content="Largest collection of Indonesian dad jokes as a consumable API" />
</svelte:head>

<section>
  <Notice emoji="💡">
    {$_('api.limit')}
  </Notice>
</section>

<section class="api_page">
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
  <p>{$_('api.get.id.body', { values: { total } })}</p>
  <Codeblock>
    GET {env.API_ENDPOINT}/id/&lcub;id&rcub;
  </Codeblock>
  <h2>{$_('api.get.total.title')}</h2>
  <p>{$_('api.get.total.body')}</p>
  <Codeblock>
    GET {env.API_ENDPOINT}/total
  </Codeblock>
</section>

<style lang="scss">
  h1 {
    @apply text-4xl;
    @apply font-bold;
    @apply py-4;
  }
  h2 {
    @apply text-2xl;
    @apply font-bold;
    @apply pt-6;
    @apply pb-1;
  }
  p { 
    @apply text-base;
    @apply opacity-80;
    @apply py-2;
  }
</style>