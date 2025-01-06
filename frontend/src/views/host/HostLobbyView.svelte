<script lang="ts">
    import Button from "../../lib/Button.svelte";
    import type { HostGame } from "../../service/host/host";

    export let game: HostGame;

    function start() {
        game.start();
    }
</script>

<div class="p-8">
    <div class="flex justify-end">
        <Button on:click={start}> Start game </Button>
    </div>
</div>