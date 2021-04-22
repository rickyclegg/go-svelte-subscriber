<svelte:options tag="my-counter"/>

<script>
  export let counter = 0

  const onClick = () => {
    counter += 1
  }

  $: counter = typeof counter === 'string' ? parseInt(counter) : counter;
</script>

<button on:click={onClick}>
    Click me!
</button>

<p>Current count: {counter}</p>

<style>
    p {
      font-family: 'Comic Sans MS';
    }
</style>