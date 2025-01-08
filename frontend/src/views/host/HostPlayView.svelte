<script lang="ts">
    import QuizChoiceCard from "../../lib/play/QuizChoiceCard.svelte";
    import { type HostGame, type tick, currentQuestion } from "../../service/host/host";
    import { COLORS } from "../../model/quiz";

    export let game: HostGame;
</script>

{#if $currentQuestion != null}
    <div class="min-h-screen h-screen flex flex-col">
        <div class="bg-white text-3x1 border-b p-4 font-bold text-center">
            {$currentQuestion.name}
        </div>
        <div class="flex-1 flex flex-col justify-center pl-4">
            <div class="flex justify-between items-center">
                <div class="bg-purple-500 text-white text-3x1 w-16 h-16 rounded-full flex items-center justify-center ml-8">
                    {$tick}
                </div>
                <img alt="center" class="max-w-[500px]"/>
                <div class="w-24"></div>
            </div>
        </div>
        <div class="flex flex-wrap w-full h-96">
            {#each COLORS as color, i}
                <QuizChoiceCard {color}>
                    <p class="pl-14"> {$currentQuestion.choices[i].name} </p>
                </QuizChoiceCard>
            {/each}
        </div>
    </div>
{/if}