<script lang="ts">
    import { PlayerGame } from "../../service/player/player";
    import PlayerJoinView from "./PlayerJoinView.svelte";

    let game = new PlayerGame();
    let active = false;

    function onJoin() {
        active = true;
    }
</script>

{#if active}
    active
{:else}
    <PlayerJoinView on:join={onJoin} {game} />
{/if}