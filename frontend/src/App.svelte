<script lang="ts">
  import { onMount } from "svelte";

  let bits: boolean[] = $state([]);

  async function getCheckboxState(): Promise<boolean[]> {
    const response = await fetch("http://localhost:8080/api/checkboxes");
    const buffer = await response.arrayBuffer();

    const byteArr = new Uint8Array(buffer);

    const bits = new Array(1_000_000);

    console.time("State generation");
    for (let i = 0; i < 125_000; i++) {
      let byte = byteArr[i];

      for (let j = 0; j < 8; j++) {
        const shift = byte >> (7 - j);

        bits[i * 8 + j] = shift % 2 === 1;
      }
    }
    console.timeEnd("State generation");

    return bits;
  }

  onMount(() => {
    getCheckboxState().then((b) => {
      bits = b;
    });
  });
</script>

<main>
  <h1>Million Checkboxes</h1>

  {#each bits as bit}
    <input type="checkbox" checked={bit} />
  {/each}
</main>
