<script lang="ts">
  import { goto } from '$app/navigation';
  import { _ } from 'svelte-i18n';
  // TODO: Create hamburger navigation (I know, just bit lazy on this one)
</script>

<nav class="flex flex-col md:flex-row py-6 font-body">
  <div on:click={() => goto('/')} class="hover:cursor-pointer flex-5 font-bold text-2xl">Jokesbapak2</div>
  <div on:click={() => goto('/why')} class="navigation-item">{$_('navigation.why')}</div>
  <div on:click={() => goto('/guide')} class="navigation-item">{$_('navigation.guide')}</div>
  <div on:click={() => goto('/api')} class="navigation-item">{$_('navigation.api')}</div>
</nav>

<style>
  .navigation-item {
    @apply flex-1 duration-300 transition ease-in-out py-2 'md:py-0' 'md:opacity-50' 'hover:cursor-pointer' 'hover:scale-110' 'hover:opacity-100';
  }
</style>
