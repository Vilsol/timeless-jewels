/// <reference types="@sveltejs/kit" />

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
declare namespace App {
  // interface Locals {}
  // interface Platform {}
  // interface Session {}
  // interface Stuff {}
}

declare global {
  function Calculate(passiveID: number, seed: number, jewelType: number, conqueror: string);

  const TimelessJewelSeedRanges: Record<number, { min: number; max: number }>;
  const TimelessJewels: Record<string, string>;
  const TimelessJewelConquerors: Record<number, string[]>;
  const TreeToPassive: Record<number, number>;
}
