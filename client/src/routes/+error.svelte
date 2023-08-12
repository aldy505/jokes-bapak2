<script context="module" lang="ts">
  import type { ErrorLoad } from '@sveltejs/kit';
  import Sentry from '$lib/logging';

  export const load: ErrorLoad = ({ error }) => {
    Sentry.captureException(error);
    return {};
  };
</script>

<script lang="ts">
  import { goto } from '$app/navigation';
  import { _ } from 'svelte-i18n';
</script>

<section>
  <h1 class="text-5xl font-bold">{$_('error.heading')}</h1>
  <p class="text-base" on:click={() => goto('/')}>{$_('error.homepage')}</p>
</section>
