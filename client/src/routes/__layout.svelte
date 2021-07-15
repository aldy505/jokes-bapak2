<script lang="ts" context="module">
  import '$lib/locale';
</script>

<script lang="ts">
  //import '../locale'
  import '@fontsource/fira-mono';
  import '@fontsource/rubik';
  import { _ } from 'svelte-i18n';
  import Navbar from '../components/navbar.svelte';
</script>

<svelte:head>
  <meta name="author" content="Reinaldy Rafli" />
  <meta name="theme-color" content="#f4a9f5" />
  <meta name="publisher" content="https://jokesbapak2.pages.dev/" />
  <meta property="og:image" content="https://jokesbapak2.pages.dev/social.jpg" />
  <meta property="og:image:alt" content="Jokesbapak2 - Largest collection of Indonesian dad jokes" />
  <meta property="og:image:secure_url" content="https://jokesbapak2.pages.dev/social.jpg" />
  <meta property="og:locale" content="en_US" />
  <meta property="og:site_name" content={$_('meta.title')} />
  <meta name="twitter:card" content="summary_large_image" />
  <meta name="twitter:creator" content="Reinaldy Rafli" />
</svelte:head>

<div
  class="bg-gradient-to-br from-transparent to-lavender-300 dark:(bg-gray-900 to-lavender-900 text-white) min-h-screen h-full w-full"
>
  <div class="container mx-auto xl:px-40 lg:px-28 md:px-20 sm:px-12 px-8 w-full">
    <header>
      <Navbar />
    </header>
    <main class="font-body">
      <slot />
    </main>

    <footer class="font-body py-8">
      <p class="text-sm opacity-50 hover:opacity-90 transition duration-300 ease-in-out">
        {$_('footer.made')} <span class="text-red-500">&#10084;</span>
        {$_('footer.indonesia')}.
        {$_('footer.available')}
        <a href="https://www.github.com/aldy505/jokes-bapak2-api" class="hover:underline">{$_('footer.github')}</a>.
      </p>
    </footer>
  </div>
</div>

<style windi:preflights:global windi:safelist:global>
</style>
