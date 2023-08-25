import IGame from "../../types/game.type";



export const gamesPlayedByPlayerByQuery = () => `
query {
  games() {
   name
  }
}
`;




