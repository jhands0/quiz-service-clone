<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import { COLORS, type QuizQuestion } from "../../model/quiz";
    import Button from "../Button.svelte";
    import QuizChoiceCard from "../play/QuizChoiceCard.svelte";

    const dispatch = createEventDispatcher();
    
    export let selectedQuestion: QuizQuestion;

    function onDelete() {
        dispatch('delete');
    }
</script>

<div class="p-14 w-full flex-1">
    <div class="flex-1 flex flex-col min-h-[95%] border">
        <div class="flex font-bold text-center bg-gray-50 items-center border-b p-2 justify-between">
            <div class="w-32"> </div>
            <input class="p-4 text-3xl bg-gray-50 text-center" bind:value={selectedQuestion.name} on:change/>
            <div class="w-32">
                <Button on:click={onDelete}> Delete </Button>
            </div>
        </div>

        <div class="flex-1 flex flex-col justify-center pl-4">
            <div class="flex justify-center items-center">
                <img src="https://placehold.co/500x250" alt="center" class="max-w-[500px]" />
            </div>
        </div>

        <div class="flex flex-wrap w-full h-96">
            {#each COLORS as color, i}
                <QuizChoiceCard {color}>
                    <div class="px-14 w-full flex gap-2">
                        <input class="rounded px-2 py-1 w-full text-black" placeholder="Choice" bind:value={selectedQuestion.choices[i].name} />
                        <input class="w-16 h-18" type="checkbox" bind:checked={selectedQuestion.choices[i].correct} />
                    </div>
                </QuizChoiceCard>
            {/each}
        </div>
    </div>
   
</div>