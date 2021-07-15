<script lang="ts" context="module">
  import type { LoadOutput } from '@sveltejs/kit';
  import Sentry from '$lib/logging';

  export function load({ error }: LoadOutput): void {
    Sentry.captureException(error);
  }
</script>

<script lang="ts">
  import { goto } from '$app/navigation';
  import { _ } from 'svelte-i18n';
</script>

<section>
  <h1 class="text-5xl font-bold">{$_('error.heading')}</h1>
  <p class="text-base" on:click={() => goto('/')}>{$_('error.homepage')}</p>
</section>
