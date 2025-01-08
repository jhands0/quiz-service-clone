<script lang="ts">
    import type { PlayerGame } from "../../service/player/player";
    import QuizChoiceCard from "../../lib/play/QuizChoiceCard.svelte";
    import { COLORS } from "../../model/quiz";

    export let game: PlayerGame;

    function onClick(i: number) {

    }
</script>

<div class="flex flex-wrap w-full min-h-screen">
    {#each COLORS as color, i}
        <QuizChoiceCard {color}>
            <button class="h-full w-full" on:click={() => onClick}>
                X
            </button>
        </QuizChoiceCard>
    {/each}
</div>