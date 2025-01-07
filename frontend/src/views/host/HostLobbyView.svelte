<script lang="ts">
    import Button from "../../lib/Button.svelte";
    import { players, type HostGame } from "../../service/host/host";
    import PlayerNameCard from "../../lib/lobby/PlayerNameCard.svelte";

    export let game: HostGame;

    function start() {
        game.start();
    }
</script>

<div class="p-8">
    <div class="flex justify-end">
        <Button on:click={start}> Start game </Button>
    </div>
    <div class="flex flex-wrap gap-2">
        {#each $players as player(player.id)}
            <PlayerNameCard {player} />
        {/each}
    </div>
</div>